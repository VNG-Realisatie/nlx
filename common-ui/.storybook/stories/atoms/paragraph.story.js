import React from 'react'
import { Flex, Box } from '@rebass/grid'

import Paragraph from 'src/Paragraph/Paragraph'

export const paragraphStory = (
    <Flex mb={3} flexWrap="wrap">
        <Box px={3} mb={4} width={[
            1/1,
            1/2,
        ]}>
            <Paragraph>
                NLX is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape with many organisations.
            </Paragraph>
        </Box>
    </Flex>
)
