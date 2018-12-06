import React, { Fragment, Component } from "react"
import { Flex, Box } from '@rebass/grid'
import Small from '../components/Small/Small'
import Title from '../components/Title/Title'
import Button from '../components/Button/Button'
import Paragraph from '../components/Paragraph/Paragraph'
import Highlight from '../components/Highlight/Highlight'

class StepFour extends Component {
	render() {
		return (
            <Fragment>
                <Box mb={4}>
                    <Small>Step 4</Small>
                    <Title>Get insight</Title>
                </Box>
                <Box mb={4}>
                    <Paragraph size="large">
						Go to <Highlight>NLX Insight</Highlight> and click an organsation to see if it exchanged your BSN and VRN.
                    </Paragraph>
                </Box>
				<Flex justifyContent="center">
					<Button variant="secondary" as="a" href="https://insight.demo.nlx.io/" target="_blank">NLX Insight</Button>
				</Flex>
            </Fragment>
		)
	}
}

export default StepFour;
