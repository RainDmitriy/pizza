import React from 'react';
import style from './Sort.module.scss';

function Sort({setSortProps, sortProps}) {

  const [isSortActive, setIsSortActive] = React.useState(false);
  const sortPropsTranslation = {
    0: "популярности",
    1: "цене",
    2: "алфавиту"
  }

  return (
    <div className={style.sort}>
              <div className={style.sortLabel}>
                <svg
                  width="10"
                  height="6"
                  viewBox="0 0 10 6"
                  fill="none"
                  style={{rotate: isSortActive ? '0deg' : '180deg'}}
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M10 5C10 5.16927 9.93815 5.31576 9.81445 5.43945C9.69075 5.56315 9.54427 5.625 9.375 5.625H0.625C0.455729 5.625 0.309245 5.56315 0.185547 5.43945C0.061849 5.31576 0 5.16927 0 5C0 4.83073 0.061849 4.68424 0.185547 4.56055L4.56055 0.185547C4.68424 0.061849 4.83073 0 5 0C5.16927 0 5.31576 0.061849 5.43945 0.185547L9.81445 4.56055C9.93815 4.68424 10 4.83073 10 5Z"
                    fill="#2C2C2C"
                  />
                </svg>
                <b>Сортировка по:</b>
                <span onClick={() => setIsSortActive(!isSortActive)}>{sortPropsTranslation[sortProps]}</span>
              </div>
              {isSortActive && <div className={style.sortPopup}>
                <ul>
                  <li className={sortProps === 0 ? style.active : ""} onClick={() => {setSortProps(0); setIsSortActive(false)}}>популярности (сначала популярные)</li>
                  <li className={sortProps === 1 ? style.active : ""} onClick={() => {setSortProps(1); setIsSortActive(false)}}>цене (сначала дешёвые)</li>
                  <li className={sortProps === 2 ? style.active : ""} onClick={() => {setSortProps(2); setIsSortActive(false)}}>алфавиту (А-Я)</li>
                </ul>
              </div>}
    </div>
  )
};

export default Sort;