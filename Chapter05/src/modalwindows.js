import React from 'react';
import CreditCardInformation from './CreditCards';

class SingInForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.state = {
            errormessage: ''
        }
    }

    handleChange(event) {
        const name = event.target.name;
        const value = event.target.value;
        this.setState({
            [name]: value
        });
    }

    handleSubmit(event) {
        event.preventDefault();
        console.log(JSON.stringify(this.state));
    }


    render() {
        let message = null;
        if (this.state.errormessage.length !== 0) {
            message = <h5 className="mb-4 text-danger">{this.state.errormessage}</h5>;

        }
        return (
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Basic Info</h5>
                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input name="email" type="email" className="form-control" id="email"  onChange={this.handleChange}/>
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Password:</label>
                        <input name="password" type="password" className="form-control" id="pass" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-success btn-large">Sign In</button>
                        </div>
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-link text-info" onClick={() => this.props.handleNewUser()}> New User? Register</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }

}

class RegisterationForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.state = {
            errormessage: ''
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        event.preventDefault();
        const name = event.target.name;
        const value = event.target.value;
        this.setState({
            [name]: value
        });
    }

    handleSubmit(event) {
        event.preventDefault();
        console.log(this.state);
    }

    render() {
        let message = null;
        if (this.state.errormessage.length !== 0) {
            message = <h5 className="mb-4 text-danger">{this.state.errormessage}</h5>;

        }
        return (
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Registeration</h5>
                    <div className="form-group">
                        <label htmlFor="username">User Name:</label>
                        <input id="username" name='username' className="form-control" placeholder='John Doe' type='text'  onChange={this.handleChange} />
                    </div>

                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input type="email" name='email' className="form-control" id="email"  onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Password:</label>
                        <input type="password" name='pass1' className="form-control" id="pass1"  onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Confirm password:</label>
                        <input type="password" name='pass2' className="form-control" id="pass2"  onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-success btn-large">Register</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}

export class SignInModalWindow extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            showRegistrationForm: false
        };
        this.handleNewUser = this.handleNewUser.bind(this);
    }

    handleNewUser() {
        this.setState({
            showRegistrationForm: true
        });
    }



    render() {
        let modalBody = <SingInForm handleNewUser={this.handleNewUser} />
        if (this.state.showRegistrationForm === true) {
            modalBody = <RegisterationForm />
        }
        return (
            <div className="modal fade" id="register" tabIndex="-1" role="dialog" aria-labelledby="Register Form" aria-hidden="true">
                <div className="modal-dialog" role="document">
                    <div className="modal-content">
                        <div className="modal-header bg-success text-white">
                            <h5 className="modal-title" id="exampleModalLabel">Sign in</h5>
                            <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div className="modal-body">
                            {modalBody}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export function BuyModalWindow(props) {
    return (
        <div className="modal fade" id="buy" tabIndex="-1" role="dialog" aria-labelledby="Register Form" aria-hidden="true">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header bg-success text-white">
                        <h5 className="modal-title">Buy Item</h5>
                        <button type="button" className="close" data-dismiss="modal" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div className="modal-body">
                        <CreditCardInformation seperator={false} show={true} operation="Charge" />
                    </div>
                </div>
            </div>
        </div>
    );
} 