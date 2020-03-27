// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import FlippingChevron from '../FlippingChevron'

// The animation uses 10% of the viewport height as an estimation for the amount of height to animate. When
// the content grows larger that this it will probably fail
export const CollapsibleWrapper = styled.div`
  cursor: pointer;

  .collapsible-enter {
    max-height: 0;
  }
  .collapsible-enter-active {
    max-height: 10vh;
    transition: max-height 300ms;
  }
  .collapsible-exit {
    max-height: 10vh;
  }
  .collapsible-exit-active {
    max-height: 0;
    transition: max-height 300ms;
  }
`

export const CollapsibleButton = styled.div`
  display: flex;
  align-items: center;
`

export const CollapsibleTitle = styled.div`
  flex-grow: 1;
`

export const CollapsibleChevron = styled(FlippingChevron)`
  width: ${(p) => p.theme.tokens.spacing06};
  height: ${(p) => p.theme.tokens.spacing06};
  flex-grow: 0;
  transition: 300ms ease-in-out;
`

export const CollapsibleBody = styled.div`
  overflow: hidden;
`
