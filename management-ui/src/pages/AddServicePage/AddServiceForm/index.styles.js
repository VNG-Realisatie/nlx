// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'
import { Field } from 'formik'

export const Form = styled.form`
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`

export const Label = styled.label`
  display: block;
  word-break: keep-all;
  margin-bottom: ${(p) => p.theme.tokens.spacing01};
  margin-top: ${(p) => p.theme.tokens.spacing06};
`

export const StyledField = styled(Field)`
  background-color: ${(p) => p.theme.colorBackgroundInput};
  display: block;
  width: 100%;
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};
  font-family: 'Source Sans Pro', sans-serif;
  padding: ${(p) => p.theme.tokens.spacing04};
  color: ${(p) => p.theme.colorTextInputLabel};
  border: 1px solid transparent;
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

  ${(p) => {
    let width
    switch (p.size) {
      case 's':
        width = '480px'
        break

      case 'm':
        width = '680px'
        break

      default:
        width = '100%'
        break
    }

    return css`
      width: ${width};
    `
  }}
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

export const StyledLabelWithInput = styled(Label)`
  display: inline-flex;
  align-items: center;
  line-height: 1rem;
  margin-right: ${(p) => p.theme.tokens.spacing08};

  input {
    margin-right: ${(p) => p.theme.tokens.spacing04};
  }
`
