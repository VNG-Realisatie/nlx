// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import LinksWrapper from '../../components/LinksWrapper'
import LinkButton from '../../components/LinkButton'
import { Container, Row, Col } from '../Grid'
import { Section, ImageCol, Image } from './index.styles'
import newsIcon from './news.svg'

const News = () => (
  <Section alternate omitArrow>
    <Container>
      <Row>
        <Col width={[1, 1, 0.6666, 0.6666]}>
          <h2>Nieuws en ontwikkelingen</h2>
          <p>
            Op de hoogte blijven van nieuws en ontwikkelingen rondom NLX?
            Regelmatig worden er (online) technische kennissessies
            georganiseerd. Iedere twee weken is de sprint review waarin de
            voortgang van het scrum team besproken wordt.
          </p>
          <LinksWrapper>
            <LinkButton
              href="https://commonground.nl/groups/view/7edd07a0-1f96-4bba-967f-ac72347f63ef/team-core-components/events"
              text="NLX event agenda"
            />
            <LinkButton href="" text="" />
          </LinksWrapper>
        </Col>

        <ImageCol width={[1, 1, 0.3333, 0.3333]}>
          <Image
            src={newsIcon}
            alt="Iemand deelt iets met behulp van een flipover"
          />
        </ImageCol>
      </Row>
    </Container>
  </Section>
)

export default News
