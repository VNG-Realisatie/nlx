// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import RemoveButton from '../RemoveButton'

// Like an h3 but without the semantic meaning of that heading
export const StyledHeading = styled.div`
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

export const StyledActionsBar = styled.div`
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding-bottom: ${(p) => p.theme.tokens.spacing05};
`

export const StyledRemoveButton = styled(RemoveButton)`
  margin-left: auto;
`

export const StyledCollapsibleBody = styled.div`
  margin-left: calc(1.25rem + ${(p) => p.theme.tokens.spacing03});
`

export const StyledCollapsibleEmptyBody = styled.p`
  margin-bottom: 0;
`
