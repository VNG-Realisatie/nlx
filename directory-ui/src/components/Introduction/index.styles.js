// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import styled from 'styled-components'
import { mediaQueries } from '@commonground/design-system'
import BaseSection, { SectionIntro } from '../../components/Section'

export const Section = styled(BaseSection)`
  padding: ${(p) => p.theme.tokens.spacing09} 0;
  background-color: #dfe5ea;
  background-image: url('/intro-bg-small.svg');

  ${mediaQueries.mdUp`
    padding: ${(p) => p.theme.tokens.spacing10} 0;
    background: url(contact/intro-bg-large.svg) center bottom no-repeat, linear-gradient(to right, rgb(230, 233, 237), rgb(205, 214, 227));
  `}
`

export const StyledSectionIntro = styled(SectionIntro)`
  /* color: ${(p) => p.theme.tokens.colorBackground}; */
`

export const Content = styled.div`
  h2 {
    margin-bottom: ${(p) => p.theme.tokens.spacing06};
  }

  a[target='_blank'] {
    position: relative;
  }

  a[target='_blank']:after {
    position: relative;
    top: 2px;
    margin-left: ${(p) => p.theme.tokens.spacing01};
    content: url('/external-link.svg');
  }
`
