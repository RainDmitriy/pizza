import React from 'react';
import style from './PizzaBlock.module.scss';
import axios from 'axios';
import { useSelector, useDispatch } from 'react-redux';
import { updateCartItems } from '../../redux/slices/cartSlice';

function PizzaBlock({ ItemId, Title, Prices, Types, Sizes, Image }) {
  const dispatch = useDispatch();
  const { cartItems } = useSelector((state) => state.cart);

  const [selectedType, setSelectedType] = React.useState(0);
  const [selectedSize, setSelectedSize] = React.useState(0);
  const [inCart, setInCart] = React.useState(
    cartItems.filter(
      (item) =>
        item.SelectedType === selectedType &&
        item.SelectedSize === selectedSize &&
        item.Title === Title,
    ),
  );

  const onClickAdd = async () => {
    try {
      if (inCart.length > 0) {
        dispatch(
          updateCartItems([
            ...cartItems.filter((item) => item.CartId !== inCart[0].CartId),
            {
              Title,
              Prices,
              Image,
              SelectedType: selectedType,
              SelectedSize: selectedSize,
              Quantity: inCart[0].Quantity + 1,
              CartId: inCart[0].CartId,
            },
          ]),
        );
        axios
          .put(`http://localhost:8080/cart/${inCart[0].CartId}`, {
            CartId: inCart[0].CartId,
            ItemId,
            SelectedSize: selectedSize,
            Quantity: inCart[0].Quantity + 1,
            SelectedType: selectedType,
            Title,
            Prices,
            Image,
          })
          .catch((e) => console.error(e));
      } else {
        dispatch(
          updateCartItems([
            ...cartItems,
            {
              Title,
              Prices,
              Image,
              SelectedType: selectedType,
              SelectedSize: selectedSize,
              Quantity: 1,
              CartId: cartItems.length > 0 ? cartItems[cartItems.length - 1].CartId + 1 : 1,
            },
          ]),
        );
        axios.post(`http://localhost:8080/cart/0`, {
          CartId: cartItems.length > 0 ? cartItems[cartItems.length - 1].CartId + 1 : 1,
          ItemId,
          SelectedSize: selectedSize,
          Quantity: 1,
          SelectedType: selectedType,
          Title,
          Prices,
          Image,
        });
      }
    } catch (e) {
      console.log('Не удалось добавить в корзину');
    }
  };

  const typeTranslate = ['тонкое', 'традиционное'];

  React.useEffect(() => {
    setInCart(
      cartItems.filter(
        (item) =>
          item.SelectedType === selectedType &&
          item.SelectedSize === selectedSize &&
          item.Title === Title,
      ),
    );
  }, [selectedType, selectedSize, cartItems]);

  return (
    <div className={style.pizzaBlock}>
      <img className={style.pizzaBlockImage} src={Image[selectedType]} alt="Pizza" />

      <h4 className={style.pizzaBlockTitle}>{Title}</h4>
      <div className={style.pizzaBlockSelector}>
        <ul style={{ marginBottom: '8px' }}>
          {Types.map((item) => (
            <li
              key={item}
              className={selectedType === item ? style.active : ''}
              onClick={() => setSelectedType(item)}>
              {typeTranslate[item]}
            </li>
          ))}
        </ul>
        <ul>
          {Sizes.map((item) => (
            <li
              key={item}
              className={selectedSize === Sizes.indexOf(item) ? style.active : ''}
              onClick={() => setSelectedSize(Sizes.indexOf(item))}>
              {item} см
            </li>
          ))}
        </ul>
      </div>
      <div className={style.pizzaBlockBottom}>
        <div className={style.pizzaBlockPrice}>{Prices[selectedSize]} ₽</div>
        <div className="button button--outline button--add" onClick={() => onClickAdd()}>
          <svg
            wIdth="12"
            height="12"
            viewBox="0 0 12 12"
            fill="none"
            xmlns="http://www.w3.org/2000/svg">
            <path
              d="M10.8 4.8H7.2V1.2C7.2 0.5373 6.6627 0 6 0C5.3373 0 4.8 0.5373 4.8 1.2V4.8H1.2C0.5373 4.8 0 5.3373 0 6C0 6.6627 0.5373 7.2 1.2 7.2H4.8V10.8C4.8 11.4627 5.3373 12 6 12C6.6627 12 7.2 11.4627 7.2 10.8V7.2H10.8C11.4627 7.2 12 6.6627 12 6C12 5.3373 11.4627 4.8 10.8 4.8Z"
              fill="white"
            />
          </svg>
          <span>Добавить</span>
          {inCart.length > 0 ? <i>{inCart[0].Quantity}</i> : ''}
        </div>
      </div>
    </div>
  );
}

export default PizzaBlock;
