package utils

import "github.com/shine-bright-team/LAAS/v2/db"

func AddInterestRateToLoad() {
	sqlQuery := `
update contracts set loan_amount = (select round(c.loan_amount + (a.interest_rate/100 * c.loan_amount ),2)
                                    from contracts c join agreements a on a.id = c.agreement_id where contracts.id = c.id  )
where contracts.id in
      ( select contracts.id from contracts join agreements a2 on a2.id = contracts.agreement_id
                            where a2.interest_rate > 0
                              and (a2.is_interest_per_month = false
                                or ( a2.is_interest_per_month = true
                                and  mod(date_part('day', now() - a2.created_at) :: int,30 )= 0 ))
                              and contracts.id
                                      not in ( select c.id from contracts c
                                          join transactions t on c.id = t.contract_id and t.is_approved = true
                                                           group by c.id, c.loan_amount
                                                           having sum(t.paid_amount) >= c.loan_amount ));
`
	db.DB.Exec(sqlQuery)

}
