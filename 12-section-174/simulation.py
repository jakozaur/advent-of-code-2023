import streamlit as st


st.title('Estimate Section 174 impact')

st.markdown("""
Businesses in the USA were able to deduct R&D as an expense in the same way as they deduct any salary or expense.

Now, they need to amortize domestic R&D over five years and foreign R&D over 15 years.
Even worse, a recent IRS ruling declares that all software development is R&D.

So now, if the company spends \$1M on R&D, they can deduct only \$200 K per year for five years.

This change is effective for tax years beginning after December 31, 2021.
It surprised many companies, as they hoped it would be repealed.

It is harrowing for bootstrapped companies as well as those on tight cash flow.
Combined with the higher cost of capital, it is a perfect storm for many software companies.
It causes additional painful cost cuts such as layoffs and impacts the competivness of US-led innovation.
""")

st.subheader("Estimate impact")

domestic_rd = st.slider("How much domestic R&D do you spend per year?", 0, 5000, 500, format="$ %dk")
foreign_rd = st.slider("How much foreign R&D do you spend per year?", 0, 5000, 500, format="$ %dk")
expenses = st.slider("What are you other expenses per year?", 0, 5000, 100, format="$ %dk")
revenue = st.slider("What is your revenue?", 0, 5000, 1000, format="$ %dk")

profit = revenue - expenses - domestic_rd - foreign_rd
new_taxable_income = revenue - expenses - domestic_rd / 5 - foreign_rd / 15
old_taxable_income = profit

if old_taxable_income < 0:
  old_taxable_income = 0
if new_taxable_income < 0:
  new_taxable_income = 0

difference = new_taxable_income - old_taxable_income

usa_tax_rate = 21
if st.checkbox("Show details"):
  st.text("Your taxable income starting from 2022: ${new_taxable_income:.1f}k".format(**locals()))
  st.text("It used to be: ${old_taxable_income:.1f}k".format(**locals()))
  st.text("So you will have to pay taxes on your new 'fake profit': ${difference:.1f}k".format(**locals()))

  usa_tax_rate = st.slider("USA corporate income tax rate", usa_tax_rate, 21 + 12, usa_tax_rate, format="%d %%")

more_tax = difference * usa_tax_rate / 100

st.markdown("Your extra tax liability for 'fake profit' in 2022 is:  **:red[${more_tax:.1f}k]**".format(**locals()))

if more_tax > 0 and profit < 0:
  st.error("Even though you are not profitable.")

  if st.checkbox("LLC partnership mode"):
    st.text("This year, LLC will distribute tax liabilities instead of distributing profits.")
    ownership = st.slider("Youw owneship share", 0, 100, 45, format="%d %%")
    tax_for_you = ownership * more_tax / 100
    st.markdown("Your extra tax liability for 'fake profit' as a shareholder:  **:red[${tax_for_you:.1f}k]**".format(**locals()))

st.markdown("""
By [Jacek Migdal](https://twitter.com/jakozaur).
""")