import React from 'react';
import style from './PizzaBlock.module.scss';
import axios from 'axios';
import { useSelector, useDispatch } from 'react-redux';
import { updateCartItems } from '../../redux/slices/cartSlice';

function PizzaBlock({ title, price, types, sizes, image }) {
  const dispatch = useDispatch();
  const { cartItems } = useSelector((state) => state.cart);

  const [selectedType, setSelectedType] = React.useState(0);
  const [selectedSize, setSelectedSize] = React.useState(0);
  const [inCart, setInCart] = React.useState(
    cartItems.filter(
      (item) =>
        item.selectedType === selectedType &&
        item.selectedSize === selectedSize &&
        item.title === title,
    ),
  );

  const onClickAdd = async () => {
    try {
      if (inCart.length > 0) {
        axios.put(`http://localhost:5000/cart/${inCart[0].id}`, {
          title,
          price,
          image,
          selectedType,
          selectedSize,
          quantity: inCart[0].quantity + 1,
        });
        dispatch(
          updateCartItems([
            ...cartItems.filter((item) => item.id !== inCart[0].id),
            {
              title,
              price,
              image,
              selectedType,
              selectedSize,
              quantity: inCart[0].quantity + 1,
              id: inCart[0].id,
            },
          ]),
        );
      } else {
        axios.post(`http://localhost:5000/cart`, {
          title,
          price,
          image,
          selectedType,
          selectedSize,
          quantity: 1,
        });
        dispatch(
          updateCartItems([
            ...cartItems,
            {
              title,
              price,
              image,
              selectedType,
              selectedSize,
              quantity: 1,
              id: cartItems.length > 0 ? cartItems[cartItems.length - 1].id + 1 : 1,
            },
          ]),
        );
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
          item.selectedType === selectedType &&
          item.selectedSize === selectedSize &&
          item.title === title,
      ),
    );
  }, [selectedType, selectedSize, cartItems]);

  return (
    <div className={style.pizzaBlock}>
      <img className={style.pizzaBlockImage} src={image[selectedType]} alt="Pizza" />

      <h4 className={style.pizzaBlockTitle}>{title}</h4>
      <div className={style.pizzaBlockSelector}>
        <ul style={{ marginBottom: '8px' }}>
          {types.map((item) => (
            <li
              key={item}
              className={selectedType === item ? style.active : ''}
              onClick={() => setSelectedType(item)}>
              {typeTranslate[item]}
            </li>
          ))}
        </ul>
        <ul>
          {sizes.map((item) => (
            <li
              key={item}
              className={selectedSize === sizes.indexOf(item) ? style.active : ''}
              onClick={() => setSelectedSize(sizes.indexOf(item))}>
              {item} см
            </li>
          ))}
        </ul>
      </div>
      <div className={style.pizzaBlockBottom}>
        <div className={style.pizzaBlockPrice}>{price[selectedSize]} ₽</div>
        <div className="button button--outline button--add" onClick={() => onClickAdd()}>
          <svg
            width="12"
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
          {inCart.length > 0 ? <i>{inCart[0].quantity}</i> : ''}
        </div>
      </div>
    </div>
  );
}

export default PizzaBlock;
