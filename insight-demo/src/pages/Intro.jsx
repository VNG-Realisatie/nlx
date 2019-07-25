// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Fragment, Component } from "react"
import { Box } from '@rebass/grid'
import Title from '../components/Title/Title'
import Paragraph from '../components/Paragraph/Paragraph'

class Intro extends Component {
	render() {
		return (
            <Fragment>
                <Box mb={4}>
                    <Title>NLX Insight demo</Title>
                </Box>
                <Box mb={4}>
                    <Paragraph size="large">
                        This demo will show how to gain insight into which organizations exchanged your personal data through NLX.
                    </Paragraph>
                </Box>
                <Paragraph size="large">
                    In this case we will demonstrate how events are logged by applying for a parking permit in Haarlem.
                </Paragraph>
            </Fragment>
		)
	}
}

export default Intro;
