// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

import { IconKey } from '../../../../icons'

export const AccessButton = styled.button`
  float: right;
  display: flex;
  align-items: center;
  visibility: hidden;
  font-weight: ${(p) => p.theme.tokens.fontWeightSemiBold};
  color: ${(p) =>
    p.disabled ? p.theme.colorTextLinkDisabled : p.theme.colorTextLink};
  background-color: transparent;
  cursor: pointer;

  tr:hover & {
    visibility: visible;
  }
`

export const StyledIconKey = styled(IconKey)`
  width: ${(p) => p.theme.tokens.fontSizeLarge};
  height: ${(p) => p.theme.tokens.fontSizeLarge};
  margin-right: ${(p) => p.theme.tokens.spacing03};
`
