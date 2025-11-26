import styles from "./MatchList.module.css";
import { useQuery } from "@apollo/client/react";
import { gql, TypedDocumentNode } from "@apollo/client";

interface Team {
  id: string;
  longName: string;
}

interface MatchType {
  id: string;
  homeTeam: Team;
  awayTeam: Team;
}

interface Competition {
  id: string;
  startMatchday: number;
}

interface CompetitionsQuery {
  listCompetitions: Competition[];
}

interface GetMatchesByMatchdayQuery {
  getMatchesByMatchday: MatchType[];
}

interface GetMatchesByMatchdayInput {
  matchday: number;
}

const GET_COMPETITIONS: TypedDocumentNode<CompetitionsQuery> = gql`
  query MyQuery {
    listCompetitions {
      id
      name
      startMatchday
    }
  }
`;

const GET_MATCHES_BY_MATCHDAY: TypedDocumentNode<
  GetMatchesByMatchdayQuery,
  GetMatchesByMatchdayInput
> = gql`
  query MyQuery($matchday: Int!) {
    getMatchesByMatchday(input: $matchday) {
      id
      homeTeam {
        id
        longName
      }
      awayTeam {
        id
        longName
      }
    }
  }
`;

interface MatchListProps {
  roundNumber: number;
  competitionNumber: number;
}

export default function MatchList({
  roundNumber,
  competitionNumber,
}: MatchListProps) {
  const { data: competitionsData, loading: competitionsLoading } =
    useQuery(GET_COMPETITIONS);

  const currentCompetition =
    competitionsData?.listCompetitions?.[competitionNumber - 1];
  const matchday = currentCompetition
    ? currentCompetition.startMatchday + roundNumber - 1
    : undefined;

  const { data: matchesData, loading: matchesLoading } = useQuery(
    GET_MATCHES_BY_MATCHDAY,
    {
      variables: { matchday: matchday! },
      skip: !matchday,
    }
  );

  const matches = matchesData?.getMatchesByMatchday || [];

  if (competitionsLoading) {
    return <div className={styles.MatchList}>Loading competitions...</div>;
  }

  if (matchesLoading) {
    return <div className={styles.MatchList}>Loading matches...</div>;
  }

  if (!currentCompetition) {
    return <div className={styles.MatchList}>Competition not found</div>;
  }

  return (
    <div className={styles.MatchList}>
      <p>Matchweek: {currentCompetition.startMatchday + roundNumber - 1}</p>
      {matches.length === 0 ? (
        <p>No matches found for this matchday</p>
      ) : (
        matches.map((match) => (
          <div key={match.id} className={styles.MatchItem}>
            <span className={`${styles.TeamName} ${styles.Home}`}>
              {match.homeTeam.longName}
            </span>
            <span>vs</span>
            <span className={`${styles.TeamName} ${styles.Away}`}>
              {match.awayTeam.longName}
            </span>
          </div>
        ))
      )}
    </div>
  );
}
