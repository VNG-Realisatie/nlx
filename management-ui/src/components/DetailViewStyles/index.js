// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

// Used as Heading above a collapsible block in *Detail views
// Like an h3 but without the semantic meaning of that heading
export const DetailHeading = styled.div`
  line-height: ${(p) => p.theme.tokens.lineHeightHeading};
  font-weight: ${(p) => p.theme.tokens.fontWeightBold};
  font-size: ${(p) => p.theme.tokens.fontSizeLarge};
  margin: 0;

  svg {
    position: relative;
    top: -${(p) => p.theme.tokens.spacing01};
    width: 1.25rem;
    height: 1.25rem;
    margin-right: ${(p) => p.theme.tokens.spacing03};
    vertical-align: text-bottom;
    fill: ${(p) => p.theme.colorTextLabel};
  }
`

export const DetailHeadingLight = styled(DetailHeading)`
  line-height: ${(p) => p.theme.tokens.lineHeightText};
  font-weight: ${(p) => p.theme.tokens.fontWeightRegular};
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};

  svg {
    vertical-align: middle;
  }
`

export const StyledCollapsibleBody = styled.div`
  margin-left: calc(1.25rem + ${(p) => p.theme.tokens.spacing03});
`

export const StyledCollapsibleEmptyBody = styled.p`
  margin-bottom: 0;
`
