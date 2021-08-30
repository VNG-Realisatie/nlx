// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { arrayOf, number } from 'prop-types'
import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'

export const Container = styled.div`
  width: 100%;
  max-width: ${(p) => p.theme.tokens.containerWidth};
  padding: 0 ${(p) => p.theme.tokens.spacing05};
  margin: 0 auto;
`

export const Row = styled.div`
  display: flex;
  flex-wrap: wrap;
  margin: 0 -${(p) => p.theme.tokens.spacing05};
`

export const Col = styled.div`
  flex: 0 1 auto;
  padding: 0 ${(p) => p.theme.tokens.spacing05};

  ${(p) => mediaQueries.xs`
    &:not(:last-child) {
      margin-bottom: ${p.width[0] === 1 ? p.theme.tokens.spacing05 : 0};
    }
  `}

  ${() => mediaQueries.mdUp`
    width: 66.6%;
  `}
`

Col.propTypes = {
  width: arrayOf(number),
}

Col.defaultProps = {
  width: [1],
}
