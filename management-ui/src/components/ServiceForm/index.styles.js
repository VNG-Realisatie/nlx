// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Checkbox } from '@commonground/design-system'

export const Form = styled.form`
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`

export const InwayCheckbox = styled(Checkbox)`
  display: block;
`

export const ServiceNameWrapper = styled.div`
  margin-top: 0;
  margin-bottom: ${(p) => p.theme.tokens.spacing10};
`
