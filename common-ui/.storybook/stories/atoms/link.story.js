import React from 'react'
import { Flex, Box } from '@rebass/grid'
import Link from 'src/Link/Link'

export const linkStory = (
    <Flex mb={3}>
        <Box mr={4}>
            <Link>
                This is a link
            </Link>
        </Box>
        <Box mr={4}>
            <Link underline>
                Underlined when hovered
            </Link>
        </Box>
    </Flex>
)
