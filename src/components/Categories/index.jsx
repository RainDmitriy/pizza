import React from 'react';
import style from './Categories.module.scss';

function Categories({activeIndex, setActiveIndex}) {
  const categories = [
    'Все',
    'Мясные',
    'Вегетарианская',
    'Гриль',
    'Острые'
  ];

  return (
    <div className={style.categories}>
      <ul>
        {categories.map((item, index) => 
          <li
            style={{marginBottom: "8px"}}
            className={activeIndex === index ? style.active : ""} 
            onClick={() => setActiveIndex(index)}
            key={item}>
            {item}
          </li>
          )
        }
      </ul>
    </div>
  )
};

export default Categories;
