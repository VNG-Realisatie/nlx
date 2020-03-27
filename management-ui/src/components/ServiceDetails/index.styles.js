// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const StyledDrawerHeading = styled.div`
  * {
    margin: 0;
  }
  margin-bottom: ${(p) => p.theme.tokens.spacing08};
`

// Like an h3 but without the semantic meaning of that heading
export const StyledHeading = styled.div`
  line-height: ${(p) => p.theme.tokens.lineHeightHeading};
  font-weight: ${(p) => p.theme.tokens.fontWeightBold};
  font-size: ${(p) => p.theme.tokens.fontSizeLarge};
  margin: 0;

  svg {
    fill: ${(p) => p.theme.colorTextLabel};
    vertical-align: text-bottom;
    width: 1.25rem;
    height: 1.25rem;
    margin-right: ${(p) => p.theme.tokens.spacing03};
    position: relative;
    top: -${(p) => p.theme.tokens.spacing01};
  }
`
export const StyledLightHeading = styled(StyledHeading)`
  line-height: ${(p) => p.theme.tokens.lineHeightText};
  font-weight: ${(p) => p.theme.tokens.fontWeightRegular};
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};

  svg {
    vertical-align: middle;
  }
`

export const StyledInwayName = styled.span`
  flex-grow: 1;
`
