import React from 'react'
import { Flex, Box } from '@rebass/grid'

import Card, {CardContent} from 'src/Card/Card'

export const cardStory = (
    <Flex mb={3}>
        <Box px={3} mb={4} width={[
            1/1,
            1/2,
        ]}>
            <Card>
                <CardContent>
                    {'This is a <Card /> with <CardContent />'}
                </CardContent>
            </Card>
        </Box>
    </Flex>
)
