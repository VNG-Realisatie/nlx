// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import LinkButton from '../../../../components/LinkButton'

export const AccessButton = styled(LinkButton)`
  float: right;
  visibility: hidden;

  tr:hover & {
    visibility: visible;
  }
`
