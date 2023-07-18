import "./scss/app.scss";
import Header from "./components/Header";
import Sort from "./components/Sort";
import Categories from "./components/Categories";
import PizzaBlock from "./components/PizzaBlock";
import React from "react";

function App() {

  const [pizzas, setPizzas] = React.useState([]);
  
  const sorted = (pizzas, sortType) => {
    switch (sortType) {
      case 0:
        return pizzas.sort((a, b) => b.rating - a.rating);
      case 1:
        return pizzas.sort((a, b) => a.price[0] - b.price[0]);
      case 2:
        return pizzas.sort((a, b) => a.title.localeCompare(b.title));
      default:
        return pizzas;
    }
  }

  const [cartItems, setCartItems] = React.useState([]);
  const [sortProps, setSortProps] = React.useState(0);

  console.log(cartItems);

  return (
      <div className="wrapper">
        <Header />
      <div className="content">
        <div className="container">
          <div className="content__top">
            <Categories />
            <Sort setSortProps={setSortProps} sortProps={sortProps} />
          </div>
          <h2 className="content__title">Все пиццы</h2>
          <div className="content__items">
            {
              sorted(pizza, sortProps).map((item) => 
                <PizzaBlock key={item.id} {...item}/>
              )
            }
          </div>
        </div>
      </div>
    </div>

  );
};

export default App;
