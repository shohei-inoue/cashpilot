import type { ReactNode } from 'react';
import styles from './Heading.module.scss';

export type HeadingLevel = 1 | 2 | 3 | 4 | 5 | 6;

type HeadingProps = {
  level: HeadingLevel;
  children: ReactNode;
  className?: string;
  id?: string;
};

const TAGS: Record<HeadingLevel, 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6'> = {
  1: 'h1',
  2: 'h2',
  3: 'h3',
  4: 'h4',
  5: 'h5',
  6: 'h6',
};

const Heading = ({ level, children, className, id }: HeadingProps) => {
  const Tag = TAGS[level];
  const levelClass = styles[`h${level}`];

  return (
    <Tag id={id} className={`${levelClass} ${className ?? ''}`.trim()}>
      {children}
    </Tag>
  );
};

export default Heading;
