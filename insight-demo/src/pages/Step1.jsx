import React, { Fragment, Component } from "react"
import { Flex, Box } from '@rebass/grid'
import Small from '../components/Small/Small'
import Title from '../components/Title/Title'
import Button from '../components/Button/Button'
import Paragraph from '../components/Paragraph/Paragraph'
import Highlight from '../components/Highlight/Highlight'

class StepOne extends Component {
	render() {
		return (
            <Fragment>
                <Box mb={4}>
                    <Small>Step 1</Small>
                    <Title>IRMA app</Title>
                </Box>
                <Box mb={3}>
                    <Paragraph size="large">
						Download and install the mobile IRMA app for <Highlight>Android</Highlight> or <Highlight>iOS</Highlight> and create an account.
                    </Paragraph>
                </Box>
				<Flex justifyContent="center" mb={4}>
                	<Box mr={3}>
						<Button variant="secondary" as="a" href="https://play.google.com/store/apps/details?id=org.irmacard.cardemu" target="_blank">Android</Button>
					</Box>
					<Button variant="secondary" as="a" href="https://itunes.apple.com/nl/app/irma-authentication/id1294092994" target="_blank">iOS</Button>
				</Flex>
                <Paragraph size="large">
                    We will use the app to acquire a demo citizen service number (BSN) and vehicle registration number (VRN).
                </Paragraph>
            </Fragment>
		)
	}
}

export default StepOne;
