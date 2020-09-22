// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled, { css } from 'styled-components'

const normalStyle = css`
  color: ${(p) => p.theme.colorTextLabel};

  &:before {
    content: '·';
    padding: 0 ${(p) => p.theme.tokens.spacing03};
  }
`

const accentedStyle = css`
  padding: 0 ${(p) => p.theme.tokens.spacing04};
  margin-left: ${(p) => p.theme.tokens.spacing04};
  color: ${(p) => p.theme.tokens.colorBackgroundAlt};
  background-color: ${(p) => p.theme.colorTextLabel};
  border-radius: ${(p) => p.theme.tokens.fontSizeMedium};
`

export const StyledAmount = styled.span`
  font-weight: ${(p) => p.theme.tokens.fontWeightRegular};

  ${(p) => (p.isAccented ? accentedStyle : normalStyle)};
`
