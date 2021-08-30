// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'
import BaseSection from '../../components/Section'
import { Col } from '../Grid'

export const Section = styled(BaseSection)`
  position: relative;
  background-image: url('news-bg-small.svg');
  z-index: -1;

  ${mediaQueries.xs`
    padding-top: 7rem;
    margin-top: ${(p) => p.theme.tokens.spacing09};
  `}

  ${mediaQueries.sm`
    padding-top: 7rem;
    margin-top: ${(p) => p.theme.tokens.spacing09};
  `}

  ${mediaQueries.mdUp`
    padding-bottom: ${(p) => p.theme.tokens.spacing11};
    background-image: url('/news-bg-large.svg');
  `}
`

export const ImageCol = styled(Col)`
  display: flex;
  justify-content: center;
  align-items: center;
`

export const Image = styled.img`
  max-width: 190px;

  ${mediaQueries.smDown`
    position: absolute;
    top: -${(p) => p.theme.tokens.spacing09};
    left: 50%;
    transform: translateX(-50%);
  `}

  ${mediaQueries.mdUp`
    width: 100%;
    max-width: 250px;
    margin: ${(p) => p.theme.tokens.spacing05} 0 -1rem;
  `}
`
