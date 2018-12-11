import React, { Fragment, Component } from "react"
import { Flex, Box } from '@rebass/grid'
import Small from '../components/Small/Small'
import Title from '../components/Title/Title'
import Button from '../components/Button/Button'
import Paragraph from '../components/Paragraph/Paragraph'
import Highlight from '../components/Highlight/Highlight'

class StepThree extends Component {
	render() {
		return (
            <Fragment>
                <Box mb={4}>
                    <Small>Step 3</Small>
                    <Title>Parking permit</Title>
                </Box>
                <Box mb={4}>
                    <Paragraph size="large">
                        Apply for a parking permit in Haarlem on our demo site. Insert the <Highlight>BSN</Highlight> and <Highlight>VRN</Highlight> you now have in your IRMA app.
                    </Paragraph>
                </Box>
				<Flex justifyContent="center">
					<Button variant="secondary" as="a" href="https://application.demo.voorbeeld-haarlem.nl/" target="_blank">Apply for parking permit</Button>
				</Flex>
            </Fragment>
		)
	}
}

export default StepThree;
