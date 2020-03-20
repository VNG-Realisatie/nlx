This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).  
See CRA docs for details on scripts and other features.

# Management UI

Dashboard application for monitoring the nlx ecosystem

## Developers

Installation: check the [README](https://gitlab.com/commonground/nlx/nlx#running-the-complete-stack-using-modd) in nlx root directory on how to install the project locally.

### Internationalisation (i18n)

We use [i18next](https://www.i18next.com/) for our multi-lingual needs together with the [react plugin](https://react.i18next.com/).

- We are using natural language keys, so that there's a sense of content in the code
- Currently we only have one namespace `common`. We may need to expand this when the app becomes bigger, or when key conflicts arise
- For using i18n in your components, you'll usually need the [useTranslation](https://react.i18next.com/latest/usetranslation-hook) hook. For nested structures, use the [Trans component](https://react.i18next.com/latest/trans-component).

> Just as a reminder:  
> Any global re-initialisation of language strings (maybe yup errors?) may have to be done as second param in i18n's init function.
