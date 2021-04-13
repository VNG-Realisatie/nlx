// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Form } from 'formik'

export const DateInputsWrapper = styled.div`
  display: flex;
`

export const DateInputWrapper = styled.div`
  margin-right: ${(p) => p.theme.tokens.spacing07};

  & > p {
    margin-bottom: 0;
  }
`

export const ButtonWrapper = styled.div`
  margin-top: ${(p) => p.theme.tokens.spacing07};
`

export const StyledForm = styled(Form)`
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`
