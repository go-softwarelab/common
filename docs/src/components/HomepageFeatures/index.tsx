import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
    title: string;
    Svg: React.ComponentType<React.ComponentProps<'svg'>>;
    description: ReactNode;
};

const FeatureList: FeatureItem[] = [
    {
        title: 'Write Go with ease',
        Svg: require('@site/static/img/easy_to_use.svg').default,
        description: (
            <>
                Simplifies everyday coding, turning complexity into clarity!
            </>
        ),
    },
    {
        title: 'Unlock the beauty of simplicity',
        Svg: require('@site/static/img/simplicity.svg').default,
        description: (
            <>
                Helps you write Go code that's as clear as a summer sky!
            </>
        ),
    },
    {
        title: 'Unleashing the firepower of iter.Seq',
        Svg: require('@site/static/img/iterators.svg').default,
        description: (
            <>
                Supercharges Go iterators (iter.Seq and iter.Seq2) for pleasant and seamless usability!"
            </>
        ),
    },
];

function Feature({title, Svg, description}: FeatureItem) {
    return (
        <div className={clsx('col col--4')}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img"/>
            </div>
            <div className="text--center padding-horiz--md">
                <Heading as="h3">{title}</Heading>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures(): ReactNode {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
