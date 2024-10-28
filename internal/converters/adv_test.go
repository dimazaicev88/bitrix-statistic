package converters

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdvBuilder_FilterAdv(t *testing.T) {
	req := require.New(t)

	t.Run("Фильтрация без where", func(t *testing.T) {
		advSQLBuilder := NewAdvConverter(filters.Filter{
			Fields:    advSelectFields,
			Skip:      10,
			Limit:     100,
			OrderBy:   "desc",
			Order:     []string{"uuid", "referer1", "benefit"},
			Operators: nil,
		})
		sql, args, err := advSQLBuilder.Convert()
		req.NoError(err)
		req.Equal(args, []any{10, 100})
		req.Equal("SELECT t1.uuid,t1.referer1,t1.referer2,t1.description,if(t2.sessions > 0, round(t2.hits / t2.sessions, 2), -1) as attent,if(t2.sessionsBack > 0, round(t2.hitsBack / t2.sessionsBack, 2), -1) as attentBack,round(t1.cost * 1.00, 2) as cost,round(t2.revenue * 1.00, 2) as revenue,round((t2.revenue - t1.cost) * 1.00, 2) as benefit,round((if(t2.sessions > 0, t1.cost / t2.sessions, null)) * 1.00, 2) as sessionCost,round((if(t2.guests > 0, t1.cost / t2.guests, null)) * 1.00, 2) as visitorCost,if(t1.cost > 0, round(((t2.revenue - t1.cost) / t1.cost) * 100, 2), -1) as roi,t2.guests,t2.newGuests,t2.favorites,t2.hosts,t2.sessions,t2.hits,t2.guestsBack,t2.favoritesBack,t2.hostsBack,t2.sessionsBack,t2.hitsBack,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = today()) \n                                                           as guestsToday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = today()) as guestsBackToday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = today()) as newGuestsToday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = today()) as favoritesToday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = today())  as favoritesBackToday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = today()) as hostsToday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = today()) as hostsBackToday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = today()) as sessionsToday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = today()) as\tsessionsBackToday,sumIf(t2.hits, toStartOfDay(t3.dateStat) = today()) as hitsToday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = today()) as hitsBackToday,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = yesterday()) as guestsYesterday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as guestsBackYesterday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = yesterday()) as newGuestsYesterday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = yesterday()) as favoritesYesterday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = yesterday()) as favoritesBackYesterday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = yesterday()) as\thostsYesterday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as hostsBackYesterday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = yesterday()) as sessionsYesterday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = yesterday()) as sessionsBackYesterday,sumIf(t3.hits, toStartOfDay(t3.dateStat) = yesterday()) as hitsYesterday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = yesterday()) as hitsBackYesterday,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas\tguestsBefYesterday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas guestsBackBefYesterday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = (yesterday() - interval\t1 day)) \n\t\t\tas newGuestsBefYesterday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas favoritesBefYesterday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = (yesterday() - interval\n\t\t\t1 day)) as favoritesBackBefYesterday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval\t1 day)) \n\t\t\tas hostsBefYesterday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas\thostsBackBefYesterday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas sessionsBefYesterday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas sessionsBackBefYesterday,sumIf(t3.hits, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas hitsBefYesterday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas hitsBackBefYesterday FROM adv t1 \n          LEFT JOIN adv_stat t2 on t1.uuid = t2.advUuid\n          LEFT JOIN adv_day t3 on t3.advUuid = t2.advUuid GROUP BY t1.cost,t1.description,t1.referer1,t1.referer2,t1.uuid,t2.favorites,t2.favoritesBack,t2.guests,t2.guestsBack,t2.hits,t2.hitsBack,t2.hosts,t2.hostsBack,t2.newGuests,t2.revenue,t2.sessions,t2.sessionsBack ORDER BY t1.uuid,t1.referer1,benefit desc LIMIT ?, ?",
			sql)
	})

	t.Run("Фильтрация с where", func(t *testing.T) {
		advSQLBuilder := NewAdvConverter(filters.Filter{
			Fields:  advSelectFields,
			Skip:    10,
			Limit:   100,
			OrderBy: "desc",
			Order:   []string{"uuid", "referer1", "benefit"},
			Operators: []filters.Operators{
				{
					Operator: "=",
					Value:    "test_ref1",
					Field:    "referer1",
				},
				{
					Operator: "or",
				},
				{
					Operator: "=",
					Value:    "test_ref2",
					Field:    "referer2",
				},
			},
		})
		sql, args, err := advSQLBuilder.Convert()
		req.NoError(err)
		req.Equal(args, []any{"test_ref1", "test_ref2", 10, 100})
		req.Equal("SELECT t1.uuid,t1.referer1,t1.referer2,t1.description,if(t2.sessions > 0, round(t2.hits / t2.sessions, 2), -1) as attent,if(t2.sessionsBack > 0, round(t2.hitsBack / t2.sessionsBack, 2), -1) as attentBack,round(t1.cost * 1.00, 2) as cost,round(t2.revenue * 1.00, 2) as revenue,round((t2.revenue - t1.cost) * 1.00, 2) as benefit,round((if(t2.sessions > 0, t1.cost / t2.sessions, null)) * 1.00, 2) as sessionCost,round((if(t2.guests > 0, t1.cost / t2.guests, null)) * 1.00, 2) as visitorCost,if(t1.cost > 0, round(((t2.revenue - t1.cost) / t1.cost) * 100, 2), -1) as roi,t2.guests,t2.newGuests,t2.favorites,t2.hosts,t2.sessions,t2.hits,t2.guestsBack,t2.favoritesBack,t2.hostsBack,t2.sessionsBack,t2.hitsBack,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = today()) \n                                                           as guestsToday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = today()) as guestsBackToday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = today()) as newGuestsToday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = today()) as favoritesToday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = today())  as favoritesBackToday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = today()) as hostsToday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = today()) as hostsBackToday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = today()) as sessionsToday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = today()) as\tsessionsBackToday,sumIf(t2.hits, toStartOfDay(t3.dateStat) = today()) as hitsToday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = today()) as hitsBackToday,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = yesterday()) as guestsYesterday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as guestsBackYesterday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = yesterday()) as newGuestsYesterday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = yesterday()) as favoritesYesterday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = yesterday()) as favoritesBackYesterday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = yesterday()) as\thostsYesterday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = yesterday()) as hostsBackYesterday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = yesterday()) as sessionsYesterday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = yesterday()) as sessionsBackYesterday,sumIf(t3.hits, toStartOfDay(t3.dateStat) = yesterday()) as hitsYesterday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = yesterday()) as hitsBackYesterday,sumIf(t3.guestsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas\tguestsBefYesterday,sumIf(t3.guestsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas guestsBackBefYesterday,sumIf(t3.newGuests, toStartOfDay(t3.dateStat) = (yesterday() - interval\t1 day)) \n\t\t\tas newGuestsBefYesterday,sumIf(t3.favorites, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas favoritesBefYesterday,sumIf(t3.favoritesBack, toStartOfDay(t3.dateStat) = (yesterday() - interval\n\t\t\t1 day)) as favoritesBackBefYesterday,sumIf(t3.hostsDay, toStartOfDay(t3.dateStat) = (yesterday() - interval\t1 day)) \n\t\t\tas hostsBefYesterday,sumIf(t3.hostsDayBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas\thostsBackBefYesterday,sumIf(t3.sessions, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas sessionsBefYesterday,sumIf(t3.sessionsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas sessionsBackBefYesterday,sumIf(t3.hits, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas hitsBefYesterday,sumIf(t3.hitsBack, toStartOfDay(t3.dateStat) = (yesterday() - interval 1 day)) \n\t\t\tas hitsBackBefYesterday FROM adv t1 \n          LEFT JOIN adv_stat t2 on t1.uuid = t2.advUuid\n          LEFT JOIN adv_day t3 on t3.advUuid = t2.advUuid WHERE t1.referer1=? AND t1.referer2=?  GROUP BY t1.cost,t1.description,t1.referer1,t1.referer2,t1.uuid,t2.favorites,t2.favoritesBack,t2.guests,t2.guestsBack,t2.hits,t2.hitsBack,t2.hosts,t2.hostsBack,t2.newGuests,t2.revenue,t2.sessions,t2.sessionsBack ORDER BY t1.uuid,t1.referer1,benefit desc LIMIT ?, ?",
			sql)
	})

}
