import React from 'react'

import { Link } from 'react-router-dom';
import { Typography } from '@material-ui/core';

const Home = () => (
	<React.Fragment>
		<Typography variant="title" color="primary" noWrap gutterBottom>
			Home 
    </Typography>
		<Typography variant="subheading" color="default" noWrap gutterBottom>
			<p>
				Welcome to the DIVA 3rd party reference implementation.
			</p>			
			<ul>
				<li>Visit the organizations using the menu items on the left.</li>
				<li>You need to sign in with your IRMA app.</li>
			</ul>				
			<p>
				<strong>
					You can only view organization logs if you have disclosed required IRMA attributes.<br />
					If you haven&apos;t, you will be asked to do so.<br />
				</strong>
			</p>			
		</Typography>
	</React.Fragment>		
);

export default Home;