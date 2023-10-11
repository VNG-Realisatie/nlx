/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

import React from 'react'
import Link from '@docusaurus/Link'
import Head from '@docusaurus/Head';

function Home() {
    return (<>
        <Head>
          <meta http-equiv="refresh" content="0; url=https://docs.fsc.nlx.io" />
        </Head>
        <p>
            Wanneer je niet automatisch wordt doorgestuurd, volg dan deze <Link to="https://docs.fsc.nlx.io">link</Link>.
        </p>
      </>
    )
}

export default Home
