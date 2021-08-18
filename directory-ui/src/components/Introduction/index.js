// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Container } from '../Grid'
import { SectionIntro } from '../../components/Section'
import { Section, Content } from './index.styles'

const Introduction = () => (
  <Section omitArrow>
    <Container>
      <SectionIntro>
        <Content>
          <h1>Directory</h1>
          <p>
            In deze NLX directory vindt u een overzicht van alle beschikbare
            API’s per NLX omgeving (demo, pre-productie en productie).
          </p>
        </Content>
      </SectionIntro>
    </Container>
  </Section>
)

export default Introduction
