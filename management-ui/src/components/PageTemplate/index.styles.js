// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Page = styled.div`
  display: flex;
  position: relative;
  align-items: flex-start;
`

export const SkipToContent = styled.a`
  background-color: ${(p) => p.theme.tokens.colorPaletteGray900};
  color: ${(p) => p.theme.tokens.colorPaletteGray100};
  position: absolute;
  top: 0;
  left: 0;
  padding: ${(p) => p.theme.tokens.spacing05};
  text-decoration: none;
  z-index: 100;

  &:not(:focus):not(:active) {
    overflow: hidden;
    width: 1px;
    height: 1px;
    clip: rect(0 0 0 0);
    clip-path: inset(50%);
    white-space: nowrap;
  }
`

export const MainWrapper = styled.div`
  flex: 1;
  height: 100%;
`

export const Main = styled.main`
  padding: ${(p) => p.theme.tokens.spacing09};
`
