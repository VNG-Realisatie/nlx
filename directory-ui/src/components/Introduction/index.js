// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Container } from '../Grid'
import { Section, StyledSectionIntro, Content } from './index.styles'

const Introduction = (props) => (
  <Section omitArrow>
    <Container>
      <StyledSectionIntro>
        <Content>
          <h1>Directory</h1>
          <p>
            In deze directory vindt u een overzicht welke organisaties, welke
            gegevensbron via NLX ontsluiten via welke API.
          </p>
        </Content>
      </StyledSectionIntro>
    </Container>
  </Section>
)

export default Introduction
