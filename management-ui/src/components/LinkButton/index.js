// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export default styled.button`
  display: flex;
  align-items: center;
  padding: 0;
  font-weight: ${(p) => p.theme.tokens.fontWeightSemiBold};
  color: ${(p) =>
    p.disabled ? p.theme.colorTextLinkDisabled : p.theme.colorTextLink};
  background-color: transparent;
  cursor: pointer;

  svg {
    width: ${(p) => p.theme.tokens.fontSizeLarge};
    height: ${(p) => p.theme.tokens.fontSizeLarge};
    margin-right: ${(p) => p.theme.tokens.spacing03};
  }
`
