import './scss/app.scss';
import React from 'react';
import axios from 'axios';
import Home from './pages/Home';
import Cart from './pages/Cart';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { loadToggle, updateItems } from './redux/slices/itemsSlice';
import { updateCartItems, updateTotalPrice } from './redux/slices/cartSlice';
import { useSelector, useDispatch } from 'react-redux';

function App() {
  const dispatch = useDispatch();
  const { cartItems } = useSelector((state) => state.cart);

  const getData = async () => {
    try {
      await axios.get('http://localhost:8080/cart/0').then((res) => {
        dispatch(updateCartItems(res.data));
      });
      await axios.get('http://localhost:8080/item/0').then((res) => {
        dispatch(updateItems(res.data));
      });
      dispatch(loadToggle(true));
    } catch (e) {
      console.log('Не удалось получить пиццы с сервера');
    }
  };

  React.useEffect(() => {
    getData();
  }, []);

  React.useEffect(() => {
    dispatch(
      updateTotalPrice(
        cartItems.reduce((sum, item) => sum + item.Prices[item.SelectedSize] * item.Quantity, 0),
      ),
    );
  }, [cartItems]);

  return (
    <Router>
      <Routes>
        <Route exact path="/" Component={Home} />
        <Route exact path="/cart" Component={Cart} />
      </Routes>
    </Router>
  );
}

export default App;
