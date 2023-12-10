import streamlit as st
import pandas as pd
import numpy as np
from data import read_housing_primary, read_salaries
import plotly.express as px
import plotly.graph_objects as go


housing_primary = read_housing_primary()
salaries = read_salaries()

# if st.checkbox('Show raw data'):
#     st.subheader('Raw data')
#     st.write(housing_primary)

# st.subheader('Price per m^2 in Polish cities')
# st.line_chart(housing_primary)


# st.bar_chart(salaries)
# st.write(salaries)

st.subheader('Current housing affordability situation')
st.markdown("I define housing affordability as the ability to pay for an apartment with a mortgage for 80% value, " + \
    "without spending more than 30% of your post-tax income on the mortgage payment.")
warszawa_idx = housing_primary.columns.tolist().index('Warszawa')
selected_city = st.selectbox('Select a city', housing_primary.columns, index=warszawa_idx)
selected_city_price_per_m2 = housing_primary[selected_city].iloc[-1]

average_salary_gross = 6000
try:
    average_salary_gross = salaries.loc[selected_city]['Wynagrodzenie brutto']
except KeyError:
    st.warning(f'Could not find average salary for {selected_city}. Using {average_salary_gross} PLN instead.')


family_income_percentage = st.slider('Family income percentage of average', 50, 500, 150, format="%d %%", step=10)
desired_apartment_size = st.slider('Desired apartment size', 40, 120, 60, format="%d m^2")


gross_to_net_percentage = 72
mortgage_years = 25
loan_to_price = 80
mortgage_real_interest_rate = 7.5

if st.checkbox("Show assumptions"):
    st.text("Average gross salary {average_salary_gross} PLN, apartment price {selected_city_price_per_m2} PLN/m^2".format(**locals()))
    st.text("Assuming mortgage of {loan_to_price}% value for {mortgage_years} years.".format(**locals()))
    st.text("Assuming {gross_to_net_percentage}% gross to take home taxation.".format(**locals()))
    st.text("Assuming {mortgage_real_interest_rate}% real mortage interest rate.".format(**locals()))
    st.text("Assuming fixed mortage payment.")
    if st.checkbox("Override assumptions"):
        gross_to_net_percentage = st.slider('Gross to net ratio', 50, 100, gross_to_net_percentage, format="%d %%")
        morgage_years = st.slider('Mortgage length in years', 15, 30, mortgage_years, format="%d years")
        loan_to_price = st.slider('Loan to price ratio', 50, 100, loan_to_price, format="%d %%")
        mortgage_real_interest_rate = st.slider('Real mortgage interest rate', 2.5, 10.0, mortgage_real_interest_rate, format="%d %%")


def calculate_fixed_mortgage_payment(mortgage_amount, mortgage_months, mortgage_real_interest_rate):
    # Convert annual rate to monthly and make it a decimal
    monthly_interest_rate = mortgage_real_interest_rate / 12 / 100

    monthly_payment = (monthly_interest_rate * mortgage_amount) / (1 - (1 + monthly_interest_rate) ** -mortgage_months)

    return monthly_payment

family_income = average_salary_gross * family_income_percentage / 100 * gross_to_net_percentage / 100
affordable_percentage = 0.3

def calculate_payment_to_income(salary_gross, price_per_m2):
    family_income = salary_gross * family_income_percentage / 100 * gross_to_net_percentage / 100
    apartment_price = price_per_m2 * desired_apartment_size
    downpayment = apartment_price * (1 - loan_to_price / 100)
    mortgage_amount = apartment_price - downpayment
    mortgage_months = mortgage_years * 12
    monthly_mortgage_payment = calculate_fixed_mortgage_payment(mortgage_amount, mortgage_months, mortgage_real_interest_rate)
    payment_to_income = monthly_mortgage_payment / family_income
    return payment_to_income, monthly_mortgage_payment, apartment_price

payment_to_income, monthly_mortgage_payment, apartment_price = \
    calculate_payment_to_income(average_salary_gross, selected_city_price_per_m2)


st.text("Family take home (post tax) income: {family_income:.2f} PLN".format(**locals()))
st.text("Apartment price: {apartment_price:.2f} PLN".format(**locals()))

st.text("Monthly payment: {monthly_mortgage_payment:.2f} PLN".format(**locals()))
st.text("Payment to income: {payment_to_income:.2%}".format(**locals()))



if payment_to_income < affordable_percentage:
    st.text("Apartment is affordable, as {payment_to_income:.2%} is not higher than {affordable_percentage:.0%}.".format(**locals()))
else:
    st.error("Apartment is not affordable, as {payment_to_income:.2%} is higher than {affordable_percentage:.0%}.".format(**locals()))
    desired_income = monthly_mortgage_payment / affordable_percentage
    desired_income_increase = desired_income / family_income - 1
    desired_income = int(desired_income)
    st.text("You would need to increase your post-tax income by {desired_income_increase:.2%} to {desired_income} PLN.".format(**locals()))

    simulation_years = 20

    st.subheader('Simulations of future apartment prices and salaries')
    st.markdown("Let's assume following price increases for the next {simulation_years} years, excluding inflation.".format(**locals()))

    annual_property_increase_percentage = st.slider('Annual property increase percentage', -5, 10, 3, format="%d %%")
    annual_salary_increase_percentage = st.slider('Annual salary increase percentage', 0, 10, 5, format="%d %%")

    # Run simulations
    forecast_data = {
        'Year': range(2023, 2023 + simulation_years),
        'Apartment price': [selected_city_price_per_m2 * (1 + annual_property_increase_percentage / 100) ** i for i in range(simulation_years)],
        'Average gross salary': [average_salary_gross * (1 + annual_salary_increase_percentage / 100) ** i for i in range(simulation_years)],
    }
    affordability_simulation = {
        'Year': range(2023, 2023 + simulation_years),
        'Affordable': [affordable_percentage * 100 for i in range(simulation_years)],
        'Mortgage as percentage of salary': [ \
            100 * calculate_payment_to_income( \
                forecast_data['Average gross salary'][i], \
                forecast_data['Apartment price'][i])[0] \
                    for i in range(simulation_years)],
    }

    when_affordable = -1
    for i in range(simulation_years):
        if affordability_simulation['Mortgage as percentage of salary'][i] < affordability_simulation['Affordable'][i]:
            when_affordable = affordability_simulation['Year'][i]
            break

    if when_affordable != -1:
        st.text("Apartmens will be affordable in {when_affordable}.".format(**locals()))
    else:
        st.warning("Apartments will not be affordable in the next {simulation_years} years.".format(**locals()))

    # Plot the simulations
    forecast_df = pd.DataFrame(forecast_data)

    fig = go.Figure()
    fig.add_trace(go.Scatter(x=forecast_df['Year'], y=forecast_df['Apartment price'], mode='lines', name='Apartment price per m^2'))
    fig.add_trace(go.Scatter(x=forecast_df['Year'], y=forecast_df['Average gross salary'], mode='lines', name='Average gross salary'))
    fig.update_layout(title='Apartments price and average salary forecast',
                  xaxis_title='Year',
                  yaxis_title='PLN')

    st.plotly_chart(fig)

    affordability_df = pd.DataFrame(affordability_simulation)

    fig2 = go.Figure()
    fig2.add_trace(go.Scatter(x=affordability_df['Year'], y=affordability_df['Affordable'], mode='lines', name='Affordable threshold', line=dict(dash='dash')))
    fig2.add_trace(go.Scatter(x=affordability_df['Year'], y=affordability_df['Mortgage as percentage of salary'], mode='lines', name='Mortgage as percentage of salary'))
    fig2.update_layout(title='Affordability simulation',
                  xaxis_title='Year',
                  yaxis_title='Percentage',
                  yaxis=dict(range=[0, min(100, max(affordability_df['Mortgage as percentage of salary']) + 10)]))
    st.plotly_chart(fig2)
    


