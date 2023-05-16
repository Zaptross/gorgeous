import react from 'react';

type Props = {
  title: string;
  theme: Pallette;
};

export function PageTitle({ title, theme }: Props) {
  return (
    <h1 style={{ color: theme.Base1 }}>
      {title}
    </h1>
  );
}