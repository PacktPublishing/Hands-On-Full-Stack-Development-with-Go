import React from 'react';
import ReactDOM from 'react-dom';

class Card extends React.Component {
    render() {
        return (
            <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
                <div className="card mb-3">
                    <img className="card-img-top" src={this.props.img} alt={this.props.imgalt} />
                    <div className="card-body">
                        <h4 className="card-title">{this.props.productname}</h4>
                        Price: <strong>{this.props.price}</strong>
                        <p className="card-text">{this.props.desc}</p>
                        <a href="#" className="btn btn-primary">Buy</a>
                    </div>
                </div>
            </div>
        );
    }
}

class CardContainer extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            cards: []
        };
    }

    componentDidMount() {
        console.log('Component Did Mount Called: ' + new Date().toLocaleString());
        fetch('cards.json')
            .then(res => res.json())
            .then((result) => {
                console.log('Fetch...');
                this.setState({
                    cards: result
                });
            });
    }

    render() {
        const cards = this.state.cards;
        let items = cards.map(
            card => <Card key={card.id} {...card} />
        );
        return (
            <div className='container pt-4'>
                <h3 className='text-center text-primary'>Products</h3>
                <div className="pt-4 row">
                    {items}
                </div>
            </div>
        );
    }
}


// ========================================

ReactDOM.render(
    <CardContainer />,
    document.getElementById('root')
);