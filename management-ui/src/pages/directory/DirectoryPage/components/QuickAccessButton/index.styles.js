// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import ButtonWithIcon from '../../../../../components/ButtonWithIcon'

export const StyledButtonWithIcon = styled(ButtonWithIcon)`
  float: right;
  visibility: hidden;

  tr:hover & {
    visibility: visible;
  }
`
