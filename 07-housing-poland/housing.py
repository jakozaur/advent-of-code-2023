import streamlit as st
import pandas as pd
import numpy as np

DATE_COLUMN = 'Date'

def quarter_to_date(raw_date):
    quarter, year = raw_date.split(' ')
    month = ""
    if quarter == 'I':
        month = '01'
    elif quarter == 'II':
        month = '04'
    elif quarter == 'III':
        month = '07'
    elif quarter == 'IV':
        month = '10'
    return pd.to_datetime(f'{year}-{month}-01')

def price_to_int(raw_price):
    raw_price = str(raw_price)
    digits_only = ''.join([char for char in raw_price if char.isdigit()])

    try:
        if digits_only == '':
            return 4000
        return int(digits_only)
    except ValueError:
        st.error(f'Could not convert {raw_price} to int')
        return 0

# Data: https://static.nbp.pl/dane/rynek-nieruchomosci/ceny_mieszkan.xlsx
housing_primary = pd.read_csv("07-housing-poland/data/housing_primary.csv", sep=";", skiprows=6, usecols=range(23, 41)) # 44 if averages
housing_primary.rename(lambda x: str(x).split('.')[0].split("*")[0], axis='columns', inplace=True)
housing_primary.rename({'Kwartał': DATE_COLUMN}, axis='columns', inplace=True)
housing_primary[DATE_COLUMN] = housing_primary[DATE_COLUMN].apply(quarter_to_date)
housing_primary.set_index(DATE_COLUMN, inplace=True)
for column in housing_primary.columns:
    if column != DATE_COLUMN:
        housing_primary[column] = housing_primary[column].apply(price_to_int)

# if st.checkbox('Show raw data'):
#     st.subheader('Raw data')
#     st.write(housing_primary)

st.subheader('Price per m^2 in Polish cities')
st.line_chart(housing_primary)

# Data: https://bdl.stat.gov.pl/bdl/dane/podgrup/tablica
# 	K40	WYNAGRODZENIA I ŚWIADCZENIA SPOŁECZNE  / 	G403	WYNAGRODZENIA 2497	Przeciętne miesięczne wynagrodzenia brutto   
salaries = pd.read_csv("07-housing-poland/data/salaries_2022.csv", sep=";", usecols=[1, 2])
salaries['Nazwa'] = salaries['Nazwa'].apply(lambda x: x.split('.')[-1].strip())
salaries.set_index('Nazwa', inplace=True)
salaries.sort_values('Nazwa', inplace=True)
salaries['Wynagrodzenie brutto'] = salaries['Wynagrodzenie brutto'].apply(lambda x: int(x.split(",")[0]))
st.bar_chart(salaries)
# st.write(salaries)

st.subheader('Affordability check')
warszawa_idx = housing_primary.columns.tolist().index('Warszawa')
selected_city = st.selectbox('Select a city', housing_primary.columns, index=warszawa_idx)
selected_city_price_per_m2 = housing_primary[selected_city].iloc[-1]

average_salary_gross = 6000
try:
    average_salary_gross = salaries.loc[selected_city]['Wynagrodzenie brutto']
except KeyError:
    st.warning(f'Could not find average salary for {selected_city}. Using {average_salary_gross} PLN instead.')


family_income_percentage = st.slider('Family income percentage of average', 50, 500, 200, format="%d %%", step=10)
desired_apartment_size = st.slider('Desired apartment size', 40, 120, 70, format="%d m^2")


gross_to_net_percentage = 72
mortgage_years = 25
loan_to_price = 80
mortgage_real_interest_rate = 7.5

if st.checkbox("Show assumptions"):
    st.text("Average gross salary {average_salary_gross} PLN, apartment price {selected_city_price_per_m2} PLN/m^2".format(**locals()))
    st.text("Assuming mortgage of {loan_to_price}% value for {mortgage_years} years.".format(**locals()))
    st.text("Assuming {gross_to_net_percentage}% gross to take home taxation.".format(**locals()))
    st.text("Assuming {mortgage_real_interest_rate}% real mortage interest rate.".format(**locals()))
    if st.checkbox("Override assumptions"):
        gross_to_net_percentage = st.slider('Gross to net ratio', 50, 100, gross_to_net_percentage, format="%d %%")
        morgage_years = st.slider('Mortgage length in years', 15, 30, mortgage_years, format="%d years")
        loan_to_price = st.slider('Loan to price ratio', 50, 100, loan_to_price, format="%d %%")
        mortgage_real_interest_rate = st.slider('Real mortgage interest rate', 2.5, 10.0, mortgage_real_interest_rate, format="%d %%")


family_income = average_salary_gross * family_income_percentage / 100 * gross_to_net_percentage / 100


apartment_price = selected_city_price_per_m2 * desired_apartment_size
downpayment = apartment_price * (1 - loan_to_price / 100)
mortgage_amount = apartment_price - downpayment
monthly_interest_rate = mortgage_real_interest_rate / 12
mortgage_months = mortgage_years * 12
monthly_mortgage_payment = mortgage_amount / mortgage_months + mortgage_amount * monthly_interest_rate / 100
payment_to_income = monthly_mortgage_payment / family_income
affordable_percentage = 0.3

st.text("Family take home (post tax) income: {family_income:.2f} PLN".format(**locals()))
st.text("Apartment price: {apartment_price:.2f} PLN".format(**locals()))

st.text("Monthly payment: {monthly_mortgage_payment:.2f} PLN".format(**locals()))
st.text("Payment to income: {payment_to_income:.2%}".format(**locals()))



if payment_to_income < affordable_percentage:
    st.text("Apartment is affordable, as {payment_to_income:.2%} is not higher than {affordable_percentage:.0%}.".format(**locals()))
else:
    st.error("Apartment is not affordable, as {payment_to_income:.2%} is higher than {affordable_percentage:.0%}.".format(**locals()))
    desired_income = monthly_mortgage_payment / affordable_percentage
    desired_income_increase = desired_income / family_income
    desired_income = int(desired_income)
    st.text("You would need to increase your post-tax income by {desired_income_increase:.2%} to {desired_income} PLN.".format(**locals()))
    


