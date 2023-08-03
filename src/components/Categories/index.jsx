import React from 'react';
import style from './Categories.module.scss';
import { useSelector, useDispatch } from 'react-redux';
import { changeFilterType } from '../../redux/slices/filterSlice';

function Categories() {
  const { filterType } = useSelector((state) => state.filter);
  const dispatch = useDispatch();

  const categories = ['Все', 'Мясные', 'Вегетарианская', 'Гриль', 'Острые'];

  return (
    <div className={style.categories}>
      <ul>
        {categories.map((item, index) => (
          <li
            style={{ marginBottom: '8px' }}
            className={filterType === index ? style.active : ''}
            onClick={() => dispatch(changeFilterType(index))}
            key={item}>
            {item}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Categories;
