import React, { Fragment, Component } from "react"
import { Flex, Box } from '@rebass/grid'
import Small from '../components/Small/Small'
import Title from '../components/Title/Title'
import Button from '../components/Button/Button'
import Paragraph from '../components/Paragraph/Paragraph'
import Highlight from '../components/Highlight/Highlight'

class StepTwo extends Component {
	render() {
		return (
            <Fragment>
                <Box mb={4}>
                    <Small>Step 2</Small>
                    <Title>IRMA attributes</Title>
                </Box>
                <Box mb={4}>
                    <Paragraph size="large">
						Click the <Highlight>Get attributes</Highlight> button and go to <Highlight>Issue VRN</Highlight> / <Highlight>Issue BSN</Highlight> to acquire them by scanning the QR code with your IRMA app.
                    </Paragraph>
                </Box>
				<Flex justifyContent="center">
					<Button variant="secondary" as="a" href="https://acc-diva-js-reference-3p.appx.cloud/" target="_blank">Get attributes</Button>
				</Flex>
            </Fragment>
		)
	}
}

export default StepTwo;
