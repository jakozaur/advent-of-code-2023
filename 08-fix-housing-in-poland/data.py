import pandas as pd


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

def read_housing_primary():
    # Data: https://static.nbp.pl/dane/rynek-nieruchomosci/ceny_mieszkan.xlsx
    housing_primary = pd.read_csv("07-housing-poland/data/housing_primary.csv", sep=";", skiprows=6, usecols=range(23, 41)) # 44 if averages
    housing_primary.rename(lambda x: str(x).split('.')[0].split("*")[0], axis='columns', inplace=True)
    housing_primary.rename({'Kwartał': DATE_COLUMN}, axis='columns', inplace=True)
    housing_primary[DATE_COLUMN] = housing_primary[DATE_COLUMN].apply(quarter_to_date)
    housing_primary.set_index(DATE_COLUMN, inplace=True)
    for column in housing_primary.columns:
        if column != DATE_COLUMN:
            housing_primary[column] = housing_primary[column].apply(price_to_int)
    return housing_primary


def read_salaries():
    # Data: https://bdl.stat.gov.pl/bdl/dane/podgrup/tablica
    # 	K40	WYNAGRODZENIA I ŚWIADCZENIA SPOŁECZNE  / 	G403	WYNAGRODZENIA 2497	Przeciętne miesięczne wynagrodzenia brutto   
    salaries = pd.read_csv("07-housing-poland/data/salaries_2022.csv", sep=";", usecols=[1, 2])
    salaries['Nazwa'] = salaries['Nazwa'].apply(lambda x: x.split('.')[-1].strip())
    salaries.set_index('Nazwa', inplace=True)
    salaries.sort_values('Nazwa', inplace=True)
    salaries['Wynagrodzenie brutto'] = salaries['Wynagrodzenie brutto'].apply(lambda x: int(x.split(",")[0]))
    return salaries


