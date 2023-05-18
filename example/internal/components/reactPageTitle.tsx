import react from 'react';
import { useTheme } from '../provider/themeProvider';

type Props = {
  title: string;
  id: string;
};

export function PageTitle({ title, id }: Props) {
  const { theme } = useTheme();

  const style = {
    color: theme.Base2,
    margin: '0 0',
    padding: '0.67em 0.67em',
  };

  return (
    <h1 style={style} id={id}>
      {title}
    </h1>
  );
}