<?php

class BaseFilter
{
    protected array $arrFilter = [];

    protected function setFilter(Operator $operator, mixed $value, string $field): void
    {
        $this->arrFilter[] = [
            'operator' => $operator->value,
            'value' => $value,
            'field' => $field
        ];
    }

    /**
     * or
     * @return $this
     */
    public function or(): static
    {
        $this->setFilter(Operator::Or, '', '');
        return $this;
    }


    /**
     * and
     * @return $this
     */
    public function and(): static
    {
        $this->setFilter(Operator::And, '', '');
        return $this;
    }

    public function build(): string
    {
        return json_encode($this->arrFilter);
    }
}