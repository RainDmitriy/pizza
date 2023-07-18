import "./scss/app.scss";
import React from "react";
import axios from "axios";
import { AppContext } from "./context";
import Home from "./pages/Home";
import Cart from "./pages/Cart";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

function App() {

  const [pizza, setPizza] = React.useState([]);
  const [cartItems, setCartItems] = React.useState([]);
  const [isLoaded, setIsLoaded] = React.useState(false);
  const [totalPrice, setTotalPrice] = React.useState(0);

  const getData = async () => {
    try {
      setIsLoaded(false);
      await axios.get("http://localhost:5000/cart").then((res) => {
        setCartItems(res.data);
      })
      await axios.get("http://localhost:5000/items").then((res) => {
        setPizza(res.data);
       }
      );
      setIsLoaded(true);
    } catch (e) {
      console.log("Не удалось получить пиццы с сервера");
    }
  }

  React.useEffect(() => {
    getData();
  }, [])

  React.useEffect(() => {
    setTotalPrice(cartItems.reduce((sum, item) => sum + item.price[item.selectedSize] * item.quantity, 0));
  }, [cartItems])

  return (
      <AppContext.Provider value={{
        pizza,
        cartItems,
        setCartItems,
        isLoaded,
        totalPrice
      }}>
        <Router>
          <Routes>
            <Route exact path="/" Component={Home}/>
            <Route exact path="/cart" Component={Cart}/>
          </Routes>
        </Router>
      </AppContext.Provider>
  );
};

export default App;
