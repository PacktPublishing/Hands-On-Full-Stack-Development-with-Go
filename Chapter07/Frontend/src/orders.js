import React from 'react';

function Order(props) {
    return (
        <div className="col-12">
            <div className="card text-center">
                <div className="card-header"><h5>{props.productname}</h5></div>
                <div className="card-body">
                    <div className="row">
                        <div className="mx-auto col-6">
                            <img src={props.small_img} alt={props.imgalt} className="img-thumbnail float-left" />
                        </div>
                        <div className="col-6">
                            <p className="card-text">{props.desc}</p>
                            <div className="mt-4">
                                Price: <strong>{props.price}</strong>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="card-footer text-muted">
                    Purchased {(new Date(props.CreatedAt)).toLocaleDateString("en-US")}
                </div>
            </div>
            <div className="mt-3" />
        </div>
    );
}

export default class OrderContainer extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            orders: []
        };
    }

    componentDidMount() {
        console.log('fetching location: ' + this.props.location);
        fetch(this.props.location)
            .then(res => res.json())
            .then((result) => {
                this.setState({
                    orders: result
                });
            });
        console.log('orders received: ' + this.state.orders);
    }

    render() {
        return (
            <div className="row mt-5">
                {this.state.orders.map(order => <Order key={order.ID} {...order} />)}
            </div>
        );
    }
}