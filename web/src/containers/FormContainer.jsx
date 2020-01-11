import React, { Component } from 'react';
import Button from '../components/Button';

class FormContainer extends Component {
	constructor(props) {
		super(props);

		this.state = {
			user: {
				username: '',
				password: ''
			}
		}
		this.handleFormSubmit = this.handleFormSubmit.bind(this);
		this.handleInput = this.handleInput.bind(this);
	}

	handleFormSubmit(e) {
		console.log("get docs");
		e.preventDefault();
		let userData = this.state.user;

		fetch(
			'https://legalsink-demo.appspot.com/docs',
			{
				method: 'GET',
			}
		).then(response => {
			console.log(response);
		});
	}

	handleGetSecret(e) {
		console.log("get resource");
		e.preventDefault();
		fetch(
			'https://legalsink-demo.appspot.com/resource',
			{
				method: 'GET',
			}
		).then(response => {
			console.log(response);
		});
	}

	handleGetBad(e) {
		console.log("NOOOOOOOOO");
		e.preventDefault();
		fetch(
			'https://legalsink-demo.appspot.com/asfeasjkfh',
			{
				method: 'GET',
			}
		).then(response => {
			console.log(response);
		});
	}

	handleInput(e) {
		let value = e.target.value;
		let name = e.target.name;
		this.setState(
			prevState => {
				return {
					user: {
						...prevState.user, [name]: value
					}
				};
			}
		);
	}
	
	render() {
		return (
			<form className="container" onSubmit={this.handleFormSubmit}>
			  <Button
				title = {'docs'}
				action = {this.handleFormSubmit}
				/>
			  <Button
				title = {'resource'}
				action = {this.handleGetSecret}
				/>
			  <Button
				title = {'!!!!!!'}
				action = {this.handleGetBad}
				/>
			</form>
		);
	}
	
}

export default FormContainer;
