import streamlit as st
import pandas as pd

import plotly.graph_objects as go

st.title('Startup financial forecast')

st.markdown("""
One of the famous Paul Graham essays was [Default Alive or Default Dead]((http://paulgraham.com/aord.html)).

Early startups usually experience burn, a euphemism for net loss. They do that in pursuit of [growth](http://www.paulgraham.com/growth.html).

Knowing if a startup in its current trajectory would survive without additional external funding is good.
""")

money_in_bank = st.slider("How much money do you have in bank?", 0, 5000, 1000, format="$ %dk")
monthly_expenses = st.slider("What are your monthly expenses?", 0, 300, 60, format="$ %dk")
monthly_revenue = st.slider("What is your monthly revenue?", 0, 300, 5, format="$ %dk")

time_to_revenue = -1
initial_revenue = -1
if monthly_revenue == 0:
    time_to_revenue = st.slider("How long will take you to get to first revenue?", 1, 24, 6, format="%d months")
    initial_revenue = st.slider("What will be your first revenue?", 1, 10, 1, format="$%d k")


what_is_your_week_to_week_growth = st.slider("What is your week to week growth?", 0.0, 10.0, 5.0, format="%.1f %%")

simulation_max_months = 12 * 10

forecast = {
    "Month": [0],
    "Expenses": [monthly_expenses],
    "Revenue": [monthly_revenue],
    "Money in bank": [money_in_bank]
}

for month in range(1, simulation_max_months):
    revenue = forecast["Revenue"][-1] * (1 + what_is_your_week_to_week_growth / 100 / 7 * 30)
    if month == time_to_revenue:
        revenue = initial_revenue
    money_in_bank = forecast["Money in bank"][-1] + revenue - monthly_expenses
    forecast["Month"].append(month)
    forecast["Revenue"].append(revenue)
    forecast["Expenses"].append(monthly_expenses)
    forecast["Money in bank"].append(money_in_bank)
    if money_in_bank < 0:
        st.error("Your startup is Default Dead in {month} months".format(**locals()))
        break
    if revenue > monthly_expenses:
        st.success("Your startup is Default Alive. Break-even in {month} months with ${money_in_bank:.1f}k left.".format(**locals()))
        break

df = pd.DataFrame(forecast)
df['Burn rate'] = df['Expenses'] - df['Revenue']

fig = go.Figure()
fig.add_trace(go.Scatter(x=df['Month'], y=df['Revenue'], mode='lines', name='Revenue'))
fig.add_trace(go.Scatter(x=df['Month'], y=df['Expenses'], mode='lines', name='Expenses'))
fig.update_layout(title='Revenue vs Expenses',
              xaxis_title='Month',
              yaxis_title='$1000 ')

st.plotly_chart(fig)

fig2 = go.Figure()
fig2.add_trace(go.Scatter(x=df['Month'], y=df['Money in bank'], mode='lines', name='Money in the bank'))
fig2.update_layout(title='Money in the bank',
              xaxis_title='Month',
              yaxis_title='$1000 ')

st.plotly_chart(fig2)
