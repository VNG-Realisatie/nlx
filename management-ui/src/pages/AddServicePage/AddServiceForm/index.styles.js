// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Field as FormikField } from 'formik'

export const Label = styled.label`
  display: block;
  word-break: keep-all;
  margin-bottom: ${(p) => p.theme.tokens.spacing01};
  margin-top: ${(p) => p.theme.tokens.spacing06};
`

export const Field = styled(FormikField)`
  background-color: ${(p) => p.theme.colorBackgroundInput};
  display: block;
  width: 100%;
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};
  font-family: 'Source Sans Pro', sans-serif;
  padding: ${(p) => p.theme.tokens.spacing04};
  color: ${(p) => p.theme.colorTextInputLabel};
  border: 1px solid ${(p) => p.theme.colorBorderInput};
  outline: none;
  line-height: ${(p) => p.theme.tokens.lineHeightText};

  &:focus {
    padding: calc(${(p) => p.theme.tokens.spacing04} - 1px);
    border: 2px solid ${(p) => p.theme.colorBorderInputFocus};
  }

  &:placeholder {
    color: ${(p) => p.theme.colorTextInputPlaceholder};
  }

  &.invalid {
    padding: calc(${(p) => p.theme.tokens.spacing04} - 1px);
    border: 2px solid ${(p) => p.theme.colorBorderInputError};
  }
`

export const Fieldset = styled.fieldset`
  border: 0 none;
  padding: 0 0 3rem 0;
`

export const Legend = styled.legend`
  line-height: ${(p) => p.theme.tokens.lineHeightHeading};
  font-weight: ${(p) => p.theme.tokens.fontWeightBold};
  font-size: ${(p) => p.theme.tokens.fontSizeLarge};
  margin: 0;
`
