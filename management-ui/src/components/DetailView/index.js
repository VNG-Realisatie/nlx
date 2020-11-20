// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export { default as SectionGroup } from './SectionGroup'

// Used as Heading above a collapsible block in *Detail views
// Like an h3 but without the semantic meaning of that heading
export const DetailHeading = styled.div`
  line-height: ${(p) => p.theme.tokens.lineHeightHeading};
  font-weight: ${(p) => p.theme.tokens.fontWeightBold};
  font-size: ${(p) => p.theme.tokens.fontSizeLarge};
  margin: 0;

  svg {
    margin-right: ${(p) => p.theme.tokens.spacing03};
    fill: ${(p) => p.theme.colorTextLabel};
  }
`

export const DetailHeadingLight = styled(DetailHeading)`
  line-height: ${(p) => p.theme.tokens.lineHeightText};
  font-weight: ${(p) => p.theme.tokens.fontWeightRegular};
  font-size: ${(p) => p.theme.tokens.fontSizeMedium};
`

export const DetailBody = styled.div`
  margin-top: ${(p) => p.theme.tokens.spacing05};
  margin-left: calc(1.25rem + ${(p) => p.theme.tokens.spacing03});
`

export const StyledCollapsibleBody = styled.div`
  margin-top: ${(p) => p.theme.tokens.spacing05};
  margin-bottom: 2px; /* For focus styling in <Table withLinks /> */
  margin-left: calc(1.25rem + ${(p) => p.theme.tokens.spacing03});
`

export const StyledCollapsibleEmptyBody = styled.p`
  margin-bottom: 0;
`
