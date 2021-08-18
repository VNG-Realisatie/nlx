// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { bool } from 'prop-types'
import styled, { css } from 'styled-components'
import { mediaQueries } from '@commonground/design-system'

export const getColor = (p, invert = false) => {
  const colors = [p.theme.tokens.colorBackground, p.theme.colorAlternateSection]
  return colors[(p.alternate + invert) % 2]
}

const arrow = css`
  &::before {
    position: absolute;
    top: 0;
    left: 50%;
    width: 0;
    height: 0;
    border-left: 32px solid transparent;
    border-right: 32px solid transparent;
    border-top: 28px solid #dfe5ea;
    margin-left: -32px;
    content: '';

    ${mediaQueries.mdUp`
      border-top-color: #dadfe7;
    `}
  }

  &::after {
    display: inline-block;
    position: absolute;
    top: 3px;
    left: 50%;
    height: 10px;
    width: 10px;
    border-color: ${(p) => p.theme.tokens.colorPaletteGray500};
    border-style: solid;
    border-width: 2px 2px 0 0;
    margin-left: -5px;
    vertical-align: top;
    transform: rotate(135deg);
    content: '';
  }
`

const Section = styled.section`
  position: relative;
  padding: ${(p) => p.theme.tokens.spacing09} 0;
  background-color: ${(p) => getColor(p)};
  background-position: center bottom;
  background-repeat: no-repeat;

  ${mediaQueries.mdUp`
    padding: ${(p) => p.theme.tokens.spacing10} 0;
  `}

  ${(p) => !p.omitArrow && arrow}
`

Section.propTypes = {
  alternate: bool,
  omitArrow: bool,
}

Section.defaultProps = {
  alternate: false,
  omitArrow: false,
}

export default Section

export const SectionIntro = styled.div`
  p {
    font-size: ${(p) => p.theme.tokens.fontSizeLarge};
    line-height: 175%;
  }
`
