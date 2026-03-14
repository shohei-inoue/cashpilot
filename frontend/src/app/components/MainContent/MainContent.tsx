import { ReactNode } from 'react';
import styles from './MainContent.module.scss';

type MainContentProps = {
  children: ReactNode;
};

const MainContent: React.FC<MainContentProps> = ({ children }) => {
  return <div className={styles.content}>{children}</div>;
};

export default MainContent;
