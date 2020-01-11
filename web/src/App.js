import React, { Component } from 'react';
import './App.css';
import FormContainer from './containers/FormContainer';

class App extends Component {
	render() {
		return (
				<div className="container">
				<h3>Reach Form</h3>
				<FormContainer />
				</div>
		);
	}
}

export default App;
