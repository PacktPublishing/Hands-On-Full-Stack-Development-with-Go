import React from 'react';
import CreditCardInformation from './CreditCards';
import cookie from 'js-cookie';
import { Modal, ModalHeader, ModalBody } from 'reactstrap';

function submitRequest(path, requestBody, handleSignedIn,handleError) {
    fetch(path, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody)
    }).then(response => response.json())
      .then(json => {
            console.log("Response received...")
            if (json.error === undefined || !json.error) {
                //save cookie if not error
                console.log("Sign in Success...");
                cookie.set("user", json);
                handleSignedIn(json);
            } else {
                handleError(json.error);
            }
        })
        .catch(error=>console.log(error));
}


class SingInForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleError = this.handleError.bind(this);
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

    handleError(error){
        this.setState({
            errormessage: error
        });
    }

    handleSubmit(event) {
        //'users/signin'
        event.preventDefault();
        submitRequest('users/signin', this.state, this.props.handleSignedIn,this.handleError);
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
                        <input name="email" type="email" className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Password:</label>
                        <input name="password" type="password" className="form-control" id="pass" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-success btn-large" >Sign In</button>
                        </div>
                        <div className="col-12 mt-2">
                            <button className="btn btn-link text-info" onClick={() => this.props.handleNewUser()}> New User? Register</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }

}

class RegistrationForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.state = {
            errormessage: ''
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleError = this.handleError.bind(this);
    }

    handleChange(event) {
        event.preventDefault();
        const name = event.target.name;
        const value = event.target.value;
        this.setState({
            [name]: value
        });
    }
    
    handleError(error){
        this.setState({
            errormessage: error
        });
    }

    handleSubmit(event) {
        event.preventDefault();
        const userInfo = this.state;
        const firstlastname = userInfo.username.split(" ");
        if (userInfo.pass1 !== userInfo.pass2) {
            alert("PASSWORDS DO NOT MATCH");
            return;
        }
        const requestBody = {
            firstname: firstlastname[0],
            lastname: firstlastname[1],
            email: userInfo.email,
            password: userInfo.pass1
        };
        submitRequest('users', requestBody, this.props.handleSignedIn,this.handleError);
        
        console.log("Registration form: " + requestBody);
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
                    <h5 className="mb-4">Registration</h5>
                    <div className="form-group">
                        <label htmlFor="username">User Name:</label>
                        <input id="username" name='username' className="form-control" placeholder='John Doe' type='text' onChange={this.handleChange} />
                    </div>

                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input type="email" name='email' className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Password:</label>
                        <input type="password" name='pass1' className="form-control" id="pass1" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="pass">Confirm password:</label>
                        <input type="password" name='pass2' className="form-control" id="pass2" onChange={this.handleChange} />
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
        this.handleModalClose = this.handleModalClose.bind(this);
    }

    handleNewUser() {
        this.setState({
          showRegistrationForm: true
        });
    }

    handleModalClose(){
        this.setState({
            showRegistrationForm: false
        });
    }
   

    render() {
        let modalBody = <SingInForm handleNewUser={this.handleNewUser} handleSignedIn={this.props.handleSignedIn} />
        if (this.state.showRegistrationForm === true) {
            modalBody = <RegistrationForm handleSignedIn={this.props.handleSignedIn} />
        }
        return (
            <Modal id="register" tabIndex="-1" role="dialog" isOpen={this.props.showModal} toggle={this.props.toggle} onClosed={this.handleModalClose}>
                <div role="document">
                    <ModalHeader toggle={this.props.toggle} className="bg-success text-white">
                        Sign in
                    </ModalHeader>
                    <ModalBody>
                        {modalBody}
                    </ModalBody>
                </div>
            </Modal>
        );
    }
}

export function BuyModalWindow(props) {
    return (
        <Modal id="buy" tabIndex="-1" role="dialog" isOpen={props.showModal} toggle={props.toggle}>
            <div role="document">
                    <ModalHeader toggle={props.toggle} className="bg-success text-white">
                        Buy Item
                    </ModalHeader>
                    <ModalBody>
                        <CreditCardInformation user={props.user} seperator={false} show={true} productid={props.productid} price={props.price} operation="Charge" toggle={props.toggle} />
                    </ModalBody>
                </div>
                      
        </Modal>
    );
} 